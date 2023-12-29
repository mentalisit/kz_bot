package config

//func Reserv(log *logger.Logger) {
//	if Instance.BotMode == "reserve" {
//		for {
//			if checkPing() {
//				fmt.Println(time.Now().Format("15:04 02-01-2006"), " ожидание")
//				time.Sleep(1 * time.Minute)
//			} else {
//				go func() {
//					go utils.UpdateRun()
//					for {
//						if checkPing() {
//							log.Info("Сервер доступен, паникую ")
//							connectRDP()
//							//time.Sleep(1 * time.Minute)
//							procs, _ := ps.Processes()
//							for _, proc := range procs {
//								if proc.Executable() == "ConsoleClient.exe" {
//									cmd := exec.Command("taskkill", "/F", "/IM", "ConsoleClient.exe")
//									_ = cmd.Run()
//								}
//							}
//							//ReservCopyPaste.WriteCloud()
//							log.Panic("dostupen")
//						}
//					}
//				}()
//				break
//			}
//		}
//	}
//}
//
//func checkPing() bool {
//	err := utils.DoWithTries(func() error {
//		if ping("google.com:80") && ping(Instance.ServerAdrr) {
//			return nil
//		} else {
//			time.Sleep(2 * time.Second)
//			if ping("google.com:80") && ping(Instance.ServerAdrr) {
//				return nil
//			} else if ping("google.com:80") && !ping(Instance.ServerAdrr) {
//				return errors.New("no ping Server Kharkov")
//			} else if !ping("google.com:80") && !ping(Instance.ServerAdrr) {
//				time.Sleep(5 * time.Second)
//				return nil
//			}
//		}
//		return nil
//	}, 10, 15*time.Second)
//	if err != nil {
//		log.Println("Error Ping DoWithTries ", err)
//		return false
//	}
//	return true
//}
//func ping(address string) bool {
//	timeout := time.Duration(3 * time.Second)
//
//	_, err := net.DialTimeout("tcp", address, timeout)
//	if err == nil {
//		return true
//	}
//	return false
//}
//
//func connectRDP() {
//	cmd := exec.Command("cmd.exe", "/C", "start", "D:\\config\\1.bat")
//	cmd.Start()
//	log.Printf("Пробуем перезапуск .")
//}
