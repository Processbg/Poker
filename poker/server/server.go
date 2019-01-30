package server

import (
    "fmt"
    "net"
    "bufio"
    "sync"
)

type HandleFunc func(*bufio.ReadWriter)

type Server struct {

    adress string
    log map[string]int
    handler map[string]HandleFunc
    listener net.Listener

    mutex sync.RWMutex
}

func (s *Server) addNewUser (userName string) error{
    
    if !s.checkIfUserExists (userName) {
        
        var password int 

        _, err := fmt.Scanf("%d", &password)
        if err != nil {
            
           return err
        }

        s.mutex.Lock()
        s.log[userName] = password
        s.mutex.Unlock()

        fmt.Println("New User: " + userName + " has been added!")
    }
    
    fmt.Println("User: " + userName + " has logged on!")
    return nil
}

func (s *Server) checkIfUserExists (userName string) bool {
    
    s.mutex.Lock()
    _, ok :=  s.log[userName]
    s.mutex.Unlock()
    
    if !ok {

        fmt.Println("User: " + userName + "does not exist! Do you want to create an account?")
        return false
    }

    return true
} 

func (s *Server) AddHandleFunction (command string, f HandleFunc){
    
    s.mutex.Lock()
    s.handler[command] = f
    s.mutex.Unlock()
} 

func (s *Server) HandleRequest (commandName string){}
