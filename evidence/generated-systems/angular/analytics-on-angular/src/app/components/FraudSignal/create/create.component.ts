import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FraudSignalService } from '../../../services/FraudSignal.service';
import { FraudSignal } from '../../../models/FraudSignal';
import { SubBaseComponent } from '../../FraudSignal/sub.base.component';

@Component({
    selector: 'app-create-fraudSignal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFraudSignalComponent extends SubBaseComponent implements OnInit {

    title = 'Add FraudSignal';

    fraudSignalForm: FormGroup;
    fraudSignal: FraudSignal;

    constructor( http: HttpClient,
        private fraudSignalService: FraudSignalService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType): void {
        this.fraudSignalService
        .addFraudSignal(name, ruleLogic, Scenario, Dataset, ModelVersion, SignalType)
            .subscribe(() => {
                this.router.navigate(['/indexFraudSignal']);
            });
    }

    ngOnInit(): void {
    }
}