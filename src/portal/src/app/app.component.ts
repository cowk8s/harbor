import { Component } from "@angular/core";
import { ThemeInterface, THEME_ARRAY } from "./services/theme";
import { clone } from "./shared/units/utils";

@Component({
    selector: 'harbor-app',
    templateUrl: 'app.component.html'
})
export class AppComponent {
    themeArray: ThemeInterface[] = clone(THEME_ARRAY);
    styleMode: string = this.themeArray[0].showStyle;
    constructor() {
        this.setTheme();
    }
    setTheme () {

    }
}