package fuzz

import "strconv"
import "github.com/tinynetwork/tinet/internal/pkg/shell"
import "github.com/tinynetwork/tinet/internal/pkg/utils"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            var test shell.NodeConfig
            test.ExecConf(content)
            return 0

        case 1:
            content := string(bytes)
            shell.HostLinkUp(content)
            return 0

        case 2:
            content := string(bytes)

            length := len(content)
            str1 := content[0:length/2]
            str2 := content[length/2:length-1]

            shell.NetnsLinkUp(str1, str2)
            return 0

        case 3:
            content := string(bytes)
            var test shell.Interface
            test.N2nLink(content)
            return 0

        case 4:
            content := string(bytes)
            var test shell.Interface
            test.S2nLink(content)
            return 0

        case 5:
            content := string(bytes)
            var test shell.Interface
            test.V2cLink(content)
            return 0

        case 6:
            content := string(bytes)
            shell.GetContainerPid(content)
            return 0

        case 7:
            var strArr = make([]string, len(bytes))
            for i, byte := range bytes {

                strArr[i] = string(byte)
            }

            utils.RemoveDuplicatesString(strArr)

        case 8:
            content := string(bytes)
            utils.Exists(content)
            return 0


        default:
            content := string(bytes)
            var test shell.Interface
            test.P2cLink(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}