import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FraudSignalService } from '../../../services/FraudSignal.service';
import { SubBaseComponent } from '../../FraudSignal/sub.base.component';


@Component({
    selector: 'app-edit-fraudSignal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFraudSignalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FraudSignal';

    fraudSignalForm: FormGroup;
    fraudSignal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FraudSignalService,
        private fb: FormBuilder
) {
        super(http);
        this.fraudSignalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      ruleLogic: ['', Validators.required],
      Scenario: ['', ],
      Dataset: ['', ],
      ModelVersion: ['', ],
      SignalType: ['', ]
        });
    }

    
    updateFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFraudSignal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFraudSignal(params['id']).subscribe(res => {
                this.fraudSignal = res;
            });
        });
    }
}