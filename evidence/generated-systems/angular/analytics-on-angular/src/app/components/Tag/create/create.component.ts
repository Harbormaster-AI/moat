import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TagService } from '../../../services/Tag.service';
import { Tag } from '../../../models/Tag';
import { SubBaseComponent } from '../../Tag/sub.base.component';

@Component({
    selector: 'app-create-tag',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTagComponent extends SubBaseComponent implements OnInit {

    title = 'Add Tag';

    tagForm: FormGroup;
    tag: Tag;

    constructor( http: HttpClient,
        private tagService: TagService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.tagForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Datasets: ['', ],
      Models: ['', ],
      ModelVersions: ['', ],
      Dashboards: ['', ],
      Reports: ['', ],
      FeatureSets: ['', ],
      Metrics: ['', ],
      Category: ['', ]
        });
    }

    
    addTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category): void {
        this.tagService
        .addTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category)
            .subscribe(() => {
                this.router.navigate(['/indexTag']);
            });
    }

    ngOnInit(): void {
    }
}