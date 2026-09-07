import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatterService } from '../../../services/Matter.service';
import { Matter } from '../../../models/Matter';
import { SubBaseComponent } from '../../Matter/sub.base.component';

@Component({
    selector: 'app-create-matter',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMatterComponent extends SubBaseComponent implements OnInit {

    title = 'Add Matter';

    matterForm: FormGroup;
    matter: Matter;

    constructor( http: HttpClient,
        private matterService: MatterService,
        private fb: FormBuilder,
        private router: Router
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

    
    addMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status): void {
        this.matterService
        .addMatter(matterName, leadCounsel, LegalHolds, Organization, DataBreaches, Contracts, MatterType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexMatter']);
            });
    }

    ngOnInit(): void {
    }
}