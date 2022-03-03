import { HttpClient } from '@angular/common/http';
import { AfterViewInit, Component, ElementRef, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { throwError as observableThrowError, forkJoin } from 'rxjs';
import { catchError } from 'rxjs/operators';
import * as SwaggerUI from 'swagger-ui';
import { DevCenterBaseDirective } from "./dev-center-base";

@Component({
  selector: 'dev-center',
  templateUrl: './dev-center.component.html',
  viewProviders: [Title],
})
export class DevCenterComponent extends DevCenterBaseDirective implements AfterViewInit, OnInit {
  private ui: any;
  constructor(
    private el: ElementRef,
    private http: HttpClient,
    public titleService: Title) {
      super(titleService);
    }

  ngAfterViewInit() {
    this.getSwaggerUI();
  }
  getSwaggerUI(){
    forkJoin([this.http.get('/images/swagger.json')])
      .pipe(catchError(error => observableThrowError(error)))
      .subscribe(jsonArr => {
        console.log(jsonArr)
      });
  }
}