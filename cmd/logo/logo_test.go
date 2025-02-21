package logo

import (
	"testing"
)

func TestLogo(t *testing.T) {

	//Testing ASCII logo
	expLOGO_ASCII := `
           __  __ _____  _____ _______         _______     ______  
     /\   |  \/  |  __ \|_   _|__   __|/\     |  __ \ \   / / __ \ 
    /  \  | \  / | |__) | | |    | |  /  \    | |__) \ \_/ / |  | |
   / /\ \ | |\/| |  _  /  | |    | | / /\ \   |  ___/ \   /| |  | |
  / ____ \| |  | | | \ \ _| |_   | |/ ____ \  | |      | | | |__| |
 /_/    \_\_|  |_|_|  \_\_____|  |_/_/    \_\ |_|      |_|  \___\_\
                                                                   
`
	if LOGO_ASCII != expLOGO_ASCII {
		t.Errorf("Expected %v, Received %v", expLOGO_ASCII, LOGO_ASCII)
	}
}
