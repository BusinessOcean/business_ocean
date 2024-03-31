package bedomain

// // TODO: Implement ExecuteDomain
// func ExecuteDomain[T IBeDomainModule](domain T) error {

// 	done := make(chan bool, 1)

// 	// listen for interrupt signal to gracefully shutdown the application
// 	go func() {
// 		sigch := make(chan os.Signal, 1)
// 		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
// 		<-sigch

// 		done <- true
// 	}()

// 	// execute the root command
// 	go func() {
// 		// note: leave to the commands to decide whether to print their error
// 		domain.Run()

// 		done <- true
// 	}()

// 	<-done

// 	// // trigger cleanups
// 	// return domain.OnTerminate().Trigger(&core.TerminateEvent{
// 	// 	App: pb,
// 	// }, func(e *core.TerminateEvent) error {
// 	// 	return e.App.ResetBootstrapState()
// 	// })

// 	return nil
// }
