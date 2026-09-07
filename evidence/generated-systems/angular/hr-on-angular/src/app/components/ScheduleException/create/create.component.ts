import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ScheduleExceptionService } from '../../../services/ScheduleException.service';
import { ScheduleException } from '../../../models/ScheduleException';
import { SubBaseComponent } from '../../ScheduleException/sub.base.component';

@Component({
    selector: 'app-create-scheduleException',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateScheduleExceptionComponent extends SubBaseComponent implements OnInit {

    title = 'Add ScheduleException';

    scheduleExceptionForm: FormGroup;
    scheduleException: ScheduleException;

    constructor( http: HttpClient,
        private scheduleExceptionService: ScheduleExceptionService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.scheduleExceptionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  date: ['', Validators.required],
      reason: ['', Validators.required],
      hours: ['', Validators.required],
      WorkSchedule: ['', ],
      Employee: ['', ]
        });
    }

    
    addScheduleException(date, reason, hours, WorkSchedule, Employee): void {
        this.scheduleExceptionService
        .addScheduleException(date, reason, hours, WorkSchedule, Employee)
            .subscribe(() => {
                this.router.navigate(['/indexScheduleException']);
            });
    }

    ngOnInit(): void {
    }
}