
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RecommendationScenarioService } from '../../../services/RecommendationScenario.service';
import { RecommendationScenario } from '../../../models/RecommendationScenario';

@Component({
    selector: 'app-index-recommendationScenario',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRecommendationScenarioComponent implements OnInit {

    recommendationScenarios: RecommendationScenario[] = [];

    constructor(
        private router: Router,
        private service: RecommendationScenarioService
) {}

    ngOnInit(): void {
        this.getRecommendationScenarios();
}

    getRecommendationScenarios(): void {
        this.service.getRecommendationScenarios().subscribe((res) => {
        this.recommendationScenarios = res;
    });
}

    deleteRecommendationScenario(id: any): void {
        this.service.deleteRecommendationScenario(id)
            .subscribe(() => {
                this.getRecommendationScenarios();
            });
    }
}