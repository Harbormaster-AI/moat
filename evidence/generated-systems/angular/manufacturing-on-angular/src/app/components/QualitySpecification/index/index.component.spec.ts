
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexQualitySpecificationComponent } from './index.component';
import { QualitySpecificationService } from '../../../services/QualitySpecification.service';

describe('IndexQualitySpecificationComponent', () => {
  let component: IndexQualitySpecificationComponent;
  let fixture: ComponentFixture<IndexQualitySpecificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexQualitySpecificationComponent
      ],
      providers: [
        QualitySpecificationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexQualitySpecificationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});