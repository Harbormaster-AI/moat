import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { Exception_Service } from '../../../services/Exception_.service';
import { SubBaseComponent } from '../../Exception_/sub.base.component';


@Component({
    selector: 'app-edit-exception_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditException_Component extends SubBaseComponent implements OnInit {

    title = 'Edit Exception_';

    exception_Form: FormGroup;
    exception_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: Exception_Service,
        private fb: FormBuilder
) {
        super(http);
        this.exception_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      justification: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      RetentionSchedule: ['', ],
      Policy: ['', ],
      Control: ['', ],
      Risk: ['', ],
      ExceptionType: ['', ],
      Status: ['', ]
        });
    }

    
    updateException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexException_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getException_(params['id']).subscribe(res => {
                this.exception_ = res;
            });
        });
    }
}