import {
    Component,
    ViewChild,
    AfterViewChecked,
    Output,
    EventEmitter,
    Input,
    OnInit,
    ChangeDetectorRef,
} from '@angular/core';
import { NgForm } from '@angular/forms';

import { User } from '../../../base/left-side-nav/user/user';

@Component({
    selector: 'new-user-form',
    templateUrl: 'new-user-form.component.html',
    styleUrls: ['./new-user-form.component.scss', '../../../common.scss']
})

export class NewUserFormComponent implements AfterViewChecked, OnInit {
    showNewPwd: boolean = false;
    showConfirmPwd: boolean = false;

    @Input() isSelfRegistration = false;
    // Notify the form value changes
    @Output() valueChange = new EventEmitter<boolean>();
    @ViewChild("newUserFrom", {static: true}) newUserForm: NgForm;
    newUser: User = new User();
    newUserFormRef: NgForm;
    confirmedPwd: string;
    timerHandler: any;
    validationStateMap: any = {};
    mailAlreadyChecked: any = {};
    userNameAlreadyChecked: any = {};
    emailTooltip = 'TOOLTIP.EMAIL';
    usernameTooltip = 'TOOLTIP.USER_NAME';
    formValueChanged = false;

    checkOnGoing: any = {};
    constructor(private session: SessionService,
        private ref: ChangeDetectorRef)

    ngOnInit() {
        this.resetState();
    }
    resetState(): void {
        this.showConfirmPwd = false;
        this.showNewPwd = false;
    }
}