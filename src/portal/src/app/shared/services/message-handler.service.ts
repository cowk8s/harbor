import { Injectable } from '@angular/core';
import { ErrorHandler } from "../units/error-handler";

@Injectable({
    providedIn: 'root',
})
export class MessageHandlerService implements ErrorHandler {
    public error(error: any): void {
        
    }
}