import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { MatterService } from '../../../services/Matter.service';
import { SubBaseComponent } from '../../Matter/sub.base.component';


@Component({
    selector: 'app-edit-matter',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditMatterComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Matter';

    matterForm: FormGroup;
    matter: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: MatterService,
        private fb: FormBuilder
) {
        super(http);
        this.matterForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  matterName: ['', Validators.required],
      leadCounsel: ['', Validators.required],
      LegalHolds: ['', ],
      Organization: ['', ],
      DataBreaches: ['', ],
      Contracts: ['', ],
      MatterType: ['', ],
      Status: ['', ]
        });
    }

    
    updateMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexMatter']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getMatter(params['id']).subscribe(res => {
                this.matter = res;
            });
        });
    }
}