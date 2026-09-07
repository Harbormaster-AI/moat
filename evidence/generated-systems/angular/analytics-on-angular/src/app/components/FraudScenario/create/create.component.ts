import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FraudScenarioService } from '../../../services/FraudScenario.service';
import { FraudScenario } from '../../../models/FraudScenario';
import { SubBaseComponent } from '../../FraudScenario/sub.base.component';

@Component({
    selector: 'app-create-fraudScenario',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFraudScenarioComponent extends SubBaseComponent implements OnInit {

    title = 'Add FraudScenario';

    fraudScenarioForm: FormGroup;
    fraudScenario: FraudScenario;

    constructor( http: HttpClient,
        private fraudScenarioService: FraudScenarioService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType): void {
        this.fraudScenarioService
        .addFraudScenario(name, riskAppetite, Models, Datasets, Alerts, Signals, DetectionType)
            .subscribe(() => {
                this.router.navigate(['/indexFraudScenario']);
            });
    }

    ngOnInit(): void {
    }
}