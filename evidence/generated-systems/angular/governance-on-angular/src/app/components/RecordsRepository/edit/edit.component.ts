import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';
import { SubBaseComponent } from '../../RecordsRepository/sub.base.component';


@Component({
    selector: 'app-edit-recordsRepository',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRecordsRepositoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit RecordsRepository';

    recordsRepositoryForm: FormGroup;
    recordsRepository: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RecordsRepositoryService,
        private fb: FormBuilder
) {
        super(http);
        this.recordsRepositoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      location: ['', Validators.required],
      ownerDepartment: ['', Validators.required],
      Organization: ['', ],
      Records: ['', ],
      Systems: ['', ],
      RetentionSchedules: ['', ],
      LegalHolds: ['', ],
      RepositoryType: ['', ]
        });
    }

    
    updateRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRecordsRepository(name, location, ownerDepartment, Organization, Records, Systems, RetentionSchedules, LegalHolds, RepositoryType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRecordsRepository']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRecordsRepository(params['id']).subscribe(res => {
                this.recordsRepository = res;
            });
        });
    }
}