
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexThirdPartyAssessmentComponent } from './index.component';
import { ThirdPartyAssessmentService } from '../../../services/ThirdPartyAssessment.service';

describe('IndexThirdPartyAssessmentComponent', () => {
  let component: IndexThirdPartyAssessmentComponent;
  let fixture: ComponentFixture<IndexThirdPartyAssessmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexThirdPartyAssessmentComponent
      ],
      providers: [
        ThirdPartyAssessmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexThirdPartyAssessmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});