import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CompetencyRatingService } from '../../../services/CompetencyRating.service';
import { CompetencyRating } from '../../../models/CompetencyRating';
import { SubBaseComponent } from '../../CompetencyRating/sub.base.component';

@Component({
    selector: 'app-create-competencyRating',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCompetencyRatingComponent extends SubBaseComponent implements OnInit {

    title = 'Add CompetencyRating';

    competencyRatingForm: FormGroup;
    competencyRating: CompetencyRating;

    constructor( http: HttpClient,
        private competencyRatingService: CompetencyRatingService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.competencyRatingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  comment: ['', Validators.required],
      Review: ['', ],
      Competency: ['', ],
      Rating: ['', ]
        });
    }

    
    addCompetencyRating(comment, Review, Competency, Rating): void {
        this.competencyRatingService
        .addCompetencyRating(comment, Review, Competency, Rating)
            .subscribe(() => {
                this.router.navigate(['/indexCompetencyRating']);
            });
    }

    ngOnInit(): void {
    }
}