import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';
import { SubBaseComponent } from '../../RecommendationScenario/sub.base.component';


@Component({
    selector: 'app-edit-recommendationScenario',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRecommendationScenarioComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RecommendationScenario';

    recommendationScenarioForm: FormGroup;
    recommendationScenario: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RecommendationScenarioService,
        private fb: FormBuilder
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

    
    updateRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRecommendationScenario(name, objective, Models, Datasets, Experiments, Alerts, RecommendationType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRecommendationScenario']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRecommendationScenario(params['id']).subscribe(res => {
                this.recommendationScenario = res;
            });
        });
    }
}