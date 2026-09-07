import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ShiftAssignmentService } from '../../../services/ShiftAssignment.service';
import { SubBaseComponent } from '../../ShiftAssignment/sub.base.component';


@Component({
    selector: 'app-edit-shiftAssignment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditShiftAssignmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ShiftAssignment';

    shiftAssignmentForm: FormGroup;
    shiftAssignment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ShiftAssignmentService,
        private fb: FormBuilder
) {
        super(http);
        this.shiftAssignmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  assignmentDate: ['', Validators.required],
      Shift: ['', ],
      Employee: ['', ],
      WorkCenter: ['', ]
        });
    }

    
    updateShiftAssignment(assignmentDate, Shift, Employee, WorkCenter): void {
        this.route.params.subscribe((params) => {

                        this.service.updateShiftAssignment(assignmentDate, Shift, Employee, WorkCenter, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexShiftAssignment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getShiftAssignment(params['id']).subscribe(res => {
                this.shiftAssignment = res;
            });
        });
    }
}