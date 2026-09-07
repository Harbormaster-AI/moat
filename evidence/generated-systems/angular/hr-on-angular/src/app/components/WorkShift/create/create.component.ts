import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { WorkShiftService } from '../../../services/WorkShift.service';
import { WorkShift } from '../../../models/WorkShift';
import { SubBaseComponent } from '../../WorkShift/sub.base.component';

@Component({
    selector: 'app-create-workShift',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateWorkShiftComponent extends SubBaseComponent implements OnInit {

    title = 'Add WorkShift';

    workShiftForm: FormGroup;
    workShift: WorkShift;

    constructor( http: HttpClient,
        private workShiftService: WorkShiftService,
        private fb: FormBuilder,
        private router: Router
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

    
    addWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek): void {
        this.workShiftService
        .addWorkShift(startTime, endTime, breakMinutes, WorkSchedule, DayOfWeek)
            .subscribe(() => {
                this.router.navigate(['/indexWorkShift']);
            });
    }

    ngOnInit(): void {
    }
}