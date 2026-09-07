
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPolicyCoverageComponent } from './index.component';
import { PolicyCoverageService } from '../../../services/PolicyCoverage.service';

describe('IndexPolicyCoverageComponent', () => {
  let component: IndexPolicyCoverageComponent;
  let fixture: ComponentFixture<IndexPolicyCoverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPolicyCoverageComponent
      ],
      providers: [
        PolicyCoverageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPolicyCoverageComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});