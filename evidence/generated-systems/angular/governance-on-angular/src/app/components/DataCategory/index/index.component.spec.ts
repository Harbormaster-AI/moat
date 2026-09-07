
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataCategoryComponent } from './index.component';
import { DataCategoryService } from '../../../services/DataCategory.service';

describe('IndexDataCategoryComponent', () => {
  let component: IndexDataCategoryComponent;
  let fixture: ComponentFixture<IndexDataCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataCategoryComponent
      ],
      providers: [
        DataCategoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataCategoryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});