import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FraudScenarioService } from '../../../services/FraudScenario.service';
import { SubBaseComponent } from '../../FraudScenario/sub.base.component';


@Component({
    selector: 'app-edit-fraudScenario',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFraudScenarioComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FraudScenario';

    fraudScenarioForm: FormGroup;
    fraudScenario: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FraudScenarioService,
        private fb: FormBuilder
) {
        super(http);
        this.fraudScenarioForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      riskAppetite: ['', Validators.required],
      Models: ['', ],
      Datasets: ['', ],
      Alerts: ['', ],
      Signals: ['', ],
      DetectionType: ['', ]
        });
    }

    
    updateFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFraudScenario']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFraudScenario(params['id']).subscribe(res => {
                this.fraudScenario = res;
            });
        });
    }
}