import { AfterViewInit, Component, Directive, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

@Directive()
export abstract class DevCenterBaseDirective implements OnInit, AfterViewInit {
  protected constructor(
    public titleService: Title) {}
  ngOnInit() {
    this.setTitle("APP")
  }
  private setTitle(key: string) {
    this.titleService.setTitle("HIII");
  }
  abstract getSwaggerUI();
  abstract ngAfterViewInit();
}