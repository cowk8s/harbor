import { Component, OnDestroy, OnInit } from "@angular/core";

const defaultLeftTime = 5;

@Component({
    selector: 'page-not-found',
    templateUrl: "not-found.component.html",
    styleUrls: ['not-found.component.scss']
})
export class PageNotFoundComponent implements OnInit, OnDestroy {
    leftSeconds: number = defaultLeftTime;
    timeInterval: any = null;

    ngOnInit(): void {
        if (!this.timeInterval) {
            this.timeInterval = setInterval(interval => {
                this.leftSeconds--;
                if (this.leftSeconds <= 0) {
                    clearInterval(this.timeInterval);
                }
            }, defaultLeftTime)
        }
    }

    ngOnDestroy(): void {
        if (this.timeInterval) {
            clearInterval(this.timeInterval);
        }
    }
}