
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexExperimentVariantComponent } from './index.component';
import { ExperimentVariantService } from '../../../services/ExperimentVariant.service';

describe('IndexExperimentVariantComponent', () => {
  let component: IndexExperimentVariantComponent;
  let fixture: ComponentFixture<IndexExperimentVariantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexExperimentVariantComponent
      ],
      providers: [
        ExperimentVariantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexExperimentVariantComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});