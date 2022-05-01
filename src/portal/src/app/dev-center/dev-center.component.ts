import { AfterViewInit, Component, OnInit } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { throwError as observableThrowError, forkJoin } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Title } from "@angular/platform-browser";
import { DevCenterBaseDirective } from "./dev-center-base";

enum SwaggerJsonUrls {
  SWAGGER1 = '/swagger.json',
  SWAGGER2 = '/swagger2.json'
}

@Component({
  selector: 'dev-center',
  templateUrl: 'dev-center.component.html',
  viewProviders: [Title],
  styleUrls: ['dev-center.component.scss']
})
export class DevCenterComponent extends DevCenterBaseDirective implements AfterViewInit, OnInit {
  constructor(
    private http: HttpClient,
    public titleService: Title) {
    super(titleService);
  }
  ngAfterViewInit() {
    this.getSwaggerUI();
  }
  getSwaggerUI() {
    forkJoin([this.http.get(SwaggerJsonUrls.SWAGGER1), this.http.get(SwaggerJsonUrls.SWAGGER2)])
      .pipe(catchError(error => observableThrowError(error)))
      .subscribe(jsonArr => {
        const json: any = {};
      })
  }
}