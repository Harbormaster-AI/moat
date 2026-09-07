import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Exception_Service } from '../../../services/Exception_.service';
import { Exception_ } from '../../../models/Exception_';
import { SubBaseComponent } from '../../Exception_/sub.base.component';

@Component({
    selector: 'app-create-exception_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateException_Component extends SubBaseComponent implements OnInit {

    title = 'Add Exception_';

    exception_Form: FormGroup;
    exception_: Exception_;

    constructor( http: HttpClient,
        private exception_Service: Exception_Service,
        private fb: FormBuilder,
        private router: Router
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

    
    addException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status): void {
        this.exception_Service
        .addException_(title, justification, startDate, endDate, RetentionSchedule, Policy, Control, Risk, ExceptionType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexException_']);
            });
    }

    ngOnInit(): void {
    }
}