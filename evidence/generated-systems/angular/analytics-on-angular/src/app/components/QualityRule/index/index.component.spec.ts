
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexQualityRuleComponent } from './index.component';
import { QualityRuleService } from '../../../services/QualityRule.service';

describe('IndexQualityRuleComponent', () => {
  let component: IndexQualityRuleComponent;
  let fixture: ComponentFixture<IndexQualityRuleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexQualityRuleComponent
      ],
      providers: [
        QualityRuleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexQualityRuleComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});