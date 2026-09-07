import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CompetencyRatingService } from '../../../services/CompetencyRating.service';
import { SubBaseComponent } from '../../CompetencyRating/sub.base.component';


@Component({
    selector: 'app-edit-competencyRating',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCompetencyRatingComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CompetencyRating';

    competencyRatingForm: FormGroup;
    competencyRating: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CompetencyRatingService,
        private fb: FormBuilder
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

    
    updateCompetencyRating(comment, Review, Competency, Rating): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCompetencyRating(comment, Review, Competency, Rating, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCompetencyRating']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCompetencyRating(params['id']).subscribe(res => {
                this.competencyRating = res;
            });
        });
    }
}