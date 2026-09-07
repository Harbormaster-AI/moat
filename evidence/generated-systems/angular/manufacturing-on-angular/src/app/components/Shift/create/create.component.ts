import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ShiftService } from '../../../services/Shift.service';
import { Shift } from '../../../models/Shift';
import { SubBaseComponent } from '../../Shift/sub.base.component';

@Component({
    selector: 'app-create-shift',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateShiftComponent extends SubBaseComponent implements OnInit {

    title = 'Add Shift';

    shiftForm: FormGroup;
    shift: Shift;

    constructor( http: HttpClient,
        private shiftService: ShiftService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.shiftForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  shiftName: ['', Validators.required],
      startTime: ['', Validators.required],
      endTime: ['', Validators.required],
      Plant: ['', ],
      Assignments: ['', ],
      ShiftType: ['', ]
        });
    }

    
    addShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType): void {
        this.shiftService
        .addShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType)
            .subscribe(() => {
                this.router.navigate(['/indexShift']);
            });
    }

    ngOnInit(): void {
    }
}