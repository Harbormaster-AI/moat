import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WorkScheduleService } from '../../../services/WorkSchedule.service';
import { WorkSchedule } from '../../../models/WorkSchedule';
import { SubBaseComponent } from '../../WorkSchedule/sub.base.component';

@Component({
    selector: 'app-create-workSchedule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWorkScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Add WorkSchedule';

    workScheduleForm: FormGroup;
    workSchedule: WorkSchedule;

    constructor( http: HttpClient,
        private workScheduleService: WorkScheduleService,
        private fb: FormBuilder,
        private router: Router
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

    
    addWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType): void {
        this.workScheduleService
        .addWorkSchedule(name, standardHoursPerWeek, Contracts, Shifts, Exceptions, ScheduleType)
            .subscribe(() => {
                this.router.navigate(['/indexWorkSchedule']);
            });
    }

    ngOnInit(): void {
    }
}