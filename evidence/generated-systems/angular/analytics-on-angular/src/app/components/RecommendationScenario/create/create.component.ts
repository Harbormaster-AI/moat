import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';
import { RecommendationScenario } from '../../../models/RecommendationScenario';
import { SubBaseComponent } from '../../RecommendationScenario/sub.base.component';

@Component({
    selector: 'app-create-recommendationScenario',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRecommendationScenarioComponent extends SubBaseComponent implements OnInit {

    title = 'Add RecommendationScenario';

    recommendationScenarioForm: FormGroup;
    recommendationScenario: RecommendationScenario;

    constructor( http: HttpClient,
        private recommendationScenarioService: RecommendationScenarioService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.recommendationScenarioForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      objective: ['', Validators.required],
      Models: ['', ],
      Datasets: ['', ],
      Experiments: ['', ],
      Alerts: ['', ],
      RecommendationType: ['', ]
        });
    }

    
    addRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType): void {
        this.recommendationScenarioService
        .addRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType)
            .subscribe(() => {
                this.router.navigate(['/indexRecommendationScenario']);
            });
    }

    ngOnInit(): void {
    }
}