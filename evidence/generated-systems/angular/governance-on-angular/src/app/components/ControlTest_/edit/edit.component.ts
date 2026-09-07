import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ControlTest_Service } from '../../../services/ControlTest_.service';
import { SubBaseComponent } from '../../ControlTest_/sub.base.component';


@Component({
    selector: 'app-edit-controlTest_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditControlTest_Component extends SubBaseComponent implements OnInit {

    title = 'Edit ControlTest_';

    controlTest_Form: FormGroup;
    controlTest_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ControlTest_Service,
        private fb: FormBuilder
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

    
    updateControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateControlTest_(name, testPeriodStart, testPeriodEnd, sampleSize, Control, Evidence, Engagement, TestType, Effectiveness, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexControlTest_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getControlTest_(params['id']).subscribe(res => {
                this.controlTest_ = res;
            });
        });
    }
}