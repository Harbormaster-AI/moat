import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ControlTest_Service } from '../../../services/ControlTest_.service';
import { ControlTest_ } from '../../../models/ControlTest_';
import { SubBaseComponent } from '../../ControlTest_/sub.base.component';

@Component({
    selector: 'app-create-controlTest_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateControlTest_Component extends SubBaseComponent implements OnInit {

    title = 'Add ControlTest_';

    controlTest_Form: FormGroup;
    controlTest_: ControlTest_;

    constructor( http: HttpClient,
        private controlTest_Service: ControlTest_Service,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.controlTest_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      testPeriodStart: ['', Validators.required],
      testPeriodEnd: ['', Validators.required],
      sampleSize: ['', Validators.required],
      Control: ['', ],
      Evidence: ['', ],
      Engagement: ['', ],
      TestType: ['', ],
      Effectiveness: ['', ],
      Status: ['', ]
        });
    }

    
    addControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status): void {
        this.controlTest_Service
        .addControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status)
            .subscribe(() => {
                this.router.navigate(['/indexControlTest_']);
            });
    }

    ngOnInit(): void {
    }
}