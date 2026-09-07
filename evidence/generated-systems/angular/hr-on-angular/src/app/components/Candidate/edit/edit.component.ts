import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CandidateService } from '../../../services/Candidate.service';
import { SubBaseComponent } from '../../Candidate/sub.base.component';


@Component({
    selector: 'app-edit-candidate',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCandidateComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Candidate';

    candidateForm: FormGroup;
    candidate: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CandidateService,
        private fb: FormBuilder
) {
        super(http);
        this.candidateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      email: ['', Validators.required],
      phone: ['', Validators.required],
      Applications: ['', ],
      Interviews: ['', ],
      Offers: ['', ],
      Documents: ['', ],
      Source: ['', ]
        });
    }

    
    updateCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCandidate']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCandidate(params['id']).subscribe(res => {
                this.candidate = res;
            });
        });
    }
}