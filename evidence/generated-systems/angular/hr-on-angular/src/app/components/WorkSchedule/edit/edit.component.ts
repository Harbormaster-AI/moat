import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { WorkScheduleService } from '../../../services/WorkSchedule.service';
import { SubBaseComponent } from '../../WorkSchedule/sub.base.component';


@Component({
    selector: 'app-edit-workSchedule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditWorkScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit WorkSchedule';

    workScheduleForm: FormGroup;
    workSchedule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: WorkScheduleService,
        private fb: FormBuilder
) {
        super(http);
        this.workScheduleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      standardHoursPerWeek: ['', Validators.required],
      Contracts: ['', ],
      Shifts: ['', ],
      Exceptions: ['', ],
      ScheduleType: ['', ]
        });
    }

    
    updateWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexWorkSchedule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getWorkSchedule(params['id']).subscribe(res => {
                this.workSchedule = res;
            });
        });
    }
}