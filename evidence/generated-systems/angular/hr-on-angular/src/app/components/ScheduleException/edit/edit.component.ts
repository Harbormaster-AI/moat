import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ScheduleExceptionService } from '../../../services/ScheduleException.service';
import { SubBaseComponent } from '../../ScheduleException/sub.base.component';


@Component({
    selector: 'app-edit-scheduleException',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditScheduleExceptionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ScheduleException';

    scheduleExceptionForm: FormGroup;
    scheduleException: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ScheduleExceptionService,
        private fb: FormBuilder
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

    
    updateScheduleException(date, reason, hours, WorkSchedule, Employee): void {
        this.route.params.subscribe((params) => {

                        this.service.updateScheduleException(date, reason, hours, WorkSchedule, Employee, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexScheduleException']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getScheduleException(params['id']).subscribe(res => {
                this.scheduleException = res;
            });
        });
    }
}