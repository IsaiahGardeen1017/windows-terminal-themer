// deno-lint-ignore-file no-fallthrough
const path = 'C://Users/IGARDEEN/AppData/Local/Packages/Microsoft.WindowsTerminal_8wekyb3d8bbwe/LocalState/settings.json'



const terminalSettings = JSON.parse(await Deno.readTextFileSync(path));
const profiles = terminalSettings.profiles.list;
const irpProfile = profiles.find((profile: any) => {
    return profile.name === 'TERM';
});

let transDef = 55;

switch (Deno.args[0]) {
    case 'f': //Forest Image
	transDef = 75;
    	irpProfile.backgroundImage = "C:\\Users\\IGARDEEN\\termscripts\\termpics\\Forest.jpg"
	irpProfile.useAcrylic = true;
	break;
    case 'zc': //Zaphorizhian Cossacks
        transDef = 75;
        irpProfile.backgroundImage = "C:\\Users\\IGARDEEN\\Pictures\\Saved Pictures\\Ilja_Jefimowitsch_Repin_-_Reply_of_the_Zaporozhian_Cossacks_-_Yorck.jpg";
        irpProfile.useAcrylic = true;
        break;
    case 'op': //Opaque
        transDef = 75;
        irpProfile.useAcrylic = false;
        irpProfile.backgroundImage = "";
        break;
    case 'ac': //Accrylic
        transDef = 25;
    default:
        irpProfile.backgroundImage = "";
        irpProfile.useAcrylic = true;
        break;
}


const transVal = Deno.args[1] ? parseInt(Deno.args[1]) : transDef;

if (transVal >= 0 && transVal <= 100) {
    irpProfile.opacity = transVal;
}


const settingsStr = JSON.stringify(terminalSettings, null, 2);

Deno.writeTextFileSync(path, settingsStr);
