import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ShiftService } from '../../../services/Shift.service';
import { SubBaseComponent } from '../../Shift/sub.base.component';


@Component({
    selector: 'app-edit-shift',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditShiftComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Shift';

    shiftForm: FormGroup;
    shift: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ShiftService,
        private fb: FormBuilder
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

    
    updateShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateShift(shiftName, startTime, endTime, Plant, Assignments, ShiftType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexShift']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getShift(params['id']).subscribe(res => {
                this.shift = res;
            });
        });
    }
}