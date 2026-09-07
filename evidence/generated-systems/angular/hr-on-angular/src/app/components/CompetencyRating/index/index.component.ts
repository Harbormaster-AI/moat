
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompetencyRatingService } from '../../../services/CompetencyRating.service';
import { CompetencyRating } from '../../../models/CompetencyRating';

@Component({
    selector: 'app-index-competencyRating',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCompetencyRatingComponent implements OnInit {

    competencyRatings: CompetencyRating[] = [];

    constructor(
        private router: Router,
        private service: CompetencyRatingService
) {}

    ngOnInit(): void {
        this.getCompetencyRatings();
}

    getCompetencyRatings(): void {
        this.service.getCompetencyRatings().subscribe((res) => {
        this.competencyRatings = res;
    });
}

    deleteCompetencyRating(id: any): void {
        this.service.deleteCompetencyRating(id)
            .subscribe(() => {
                this.getCompetencyRatings();
            });
    }
}