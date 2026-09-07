
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexContentCategoryComponent } from './index.component';
import { ContentCategoryService } from '../../../services/ContentCategory.service';

describe('IndexContentCategoryComponent', () => {
  let component: IndexContentCategoryComponent;
  let fixture: ComponentFixture<IndexContentCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexContentCategoryComponent
      ],
      providers: [
        ContentCategoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexContentCategoryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});