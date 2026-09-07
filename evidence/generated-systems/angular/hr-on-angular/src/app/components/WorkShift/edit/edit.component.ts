import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WorkShiftService } from '../../../services/WorkShift.service';
import { SubBaseComponent } from '../../WorkShift/sub.base.component';


@Component({
    selector: 'app-edit-workShift',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWorkShiftComponent extends SubBaseComponent implements OnInit {

    title = 'Edit WorkShift';

    workShiftForm: FormGroup;
    workShift: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WorkShiftService,
        private fb: FormBuilder
) {
        super(http);
        this.workShiftForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  startTime: ['', Validators.required],
      endTime: ['', Validators.required],
      breakMinutes: ['', Validators.required],
      WorkSchedule: ['', ],
      DayOfWeek: ['', ]
        });
    }

    
    updateWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWorkShift']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWorkShift(params['id']).subscribe(res => {
                this.workShift = res;
            });
        });
    }
}