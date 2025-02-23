package logo

import (
	"testing"
)

const wantLogo = `
           __  __ _____  _____ _______         _______     ______  
     /\   |  \/  |  __ \|_   _|__   __|/\     |  __ \ \   / / __ \ 
    /  \  | \  / | |__) | | |    | |  /  \    | |__) \ \_/ / |  | |
   / /\ \ | |\/| |  _  /  | |    | | / /\ \   |  ___/ \   /| |  | |
  / ____ \| |  | | | \ \ _| |_   | |/ ____ \  | |      | | | |__| |
 /_/    \_\_|  |_|_|  \_\_____|  |_/_/    \_\ |_|      |_|  \___\_\
                                                                   
`

func TestLogo(t *testing.T) {
	t.Run("TestLogoASCII", func(t *testing.T) {
		t.Parallel()
		if LOGO_ASCII != wantLogo {
			t.Errorf("Expected %v, Received %v", wantLogo, LOGO_ASCII)
		}
	})
}
