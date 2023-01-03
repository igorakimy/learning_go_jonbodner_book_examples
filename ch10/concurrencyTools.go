package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type Input struct {
	A AIn
	B BIn
	C CIn
}

type AOut struct{}
type AIn struct{}

type BOut struct{}
type BIn struct{}

type COut struct{}
type CIn struct {
	A AOut
	B BOut
}

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()
	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()
	// Функция getResultC вызывается лишь в том случае, когда функции
	// getResultB, getResultA успешно выполняются в течение 50 миллисекунд.
	go func() {
		select {
		// Эта ветвь срабатывает в случае отмены контекста.
		case <-ctx.Done():
			return
		// Эта ветвь срабатывает при наличии данных, необходимых для вызова
		// функции getResultC.
		case inputC := <-p.inC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0
	for count < 2 {
		select {
		// Первые две ветви производят чтение из каналов, в которые делают запись
		// первые две горутины, и заполняют поля экземпляра InputC.
		case a := <-p.outA:
			inputC.A = a
			count++
		case b := <-p.outB:
			inputC.B = b
			count++
		// Следующие две ветки служат для обработки состояния ошибки.
		// Если была записана ошибка в канал p.errs, то мы возвращаем эту ошибку.
		case err := <-p.errs:
			return CIn{}, err
		// В случае отмены контекста мы возвращаем ошибку, которая сообщает об
		// отмене запроса.
		case <-ctx.Done():
			return CIn{}, ctx.Err()
		}
	}
	// Если выполняются первые две ветки, то мы выходим из цикла for-select
	// и возвращаем значение экземпляра inputC и ошибку, равную nil.
	return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	// В случае успешного выполнения функции getResultC считываем ее результат из канала
	// p.outC и возвращаем его.
	case out := <-p.outC:
		return out, nil
	// Если функция getResultC возвратила ошибку, считываем эту ошибку из канала
	// p.errs и возвращаем ее.
	case err := <-p.errs:
		return COut{}, err
	// При отмене контекста возвращаем ошибку, которая сообщает об этом.
	case <-ctx.Done():
		return COut{}, ctx.Err()
	}
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	// Определить контекст с таймаутом в 50 миллисекунд.
	// По истечении тайм-аута производится отмена контекста.
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	// Гарантированный отложенный вызов функции отмены контекста cancel
	// во избежание утечки ресурсов.
	defer cancel()
	// Заполнить экземпляр структуры processor набором каналов, которые будут
	// использованы для обмена данными с горутинами.
	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan CIn, 1),
		outC: make(chan COut, 1),
		// Размер буфера равен 2, т.к. в него потенциально могут быть записаны две ошибки.
		errs: make(chan error, 2),
	}
	// Вызов метода launch, чтобы запустить 3 горутины, которые производят вызов функций
	// getResultA, getResultB, getResultC
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}
	// Если всё в порядке, тогда записываем значение переменной inputC в канал p.inC.
	p.inC <- inputC
	// И вызываем метод waitForC в структуре processor.
	out, err := p.waitForC(ctx)
	// Вернуть результат вызывающей стороне.
	return out, err
}

func getResultA(ctx context.Context, in AIn) (AOut, error) {
	return AOut{}, nil
}

func getResultB(ctx context.Context, in BIn) (BOut, error) {
	return BOut{}, nil
}

func getResultC(ctx context.Context, in CIn) (COut, error) {
	return COut{}, nil
}

func main() {
	ctx := context.Background()
	input := Input{
		AIn{},
		BIn{},
		CIn{
			A: AOut{},
			B: BOut{},
		},
	}
	out, err := GatherAndProcess(ctx, input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reflect.TypeOf(out))
}
