import { BrowserModule } from '@angular/platform-browser';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { AppComponent } from './app.component';
import { HarborRoutingModule } from './harbor-routing.module';
import { ErrorHandler } from "./shared/units/error-handler";
import { MessageHandlerService } from "./shared/services/message-handler.service";

@NgModule({
    declarations: [
        AppComponent,
    ],
    imports: [
        BrowserModule,
        HarborRoutingModule,
    ],
    providers: [
        { provide: ErrorHandler, useClass: MessageHandlerService }
    ],
    schemas: [
        CUSTOM_ELEMENTS_SCHEMA
    ],
    bootstrap: [AppComponent]
})
export class AppModule { }