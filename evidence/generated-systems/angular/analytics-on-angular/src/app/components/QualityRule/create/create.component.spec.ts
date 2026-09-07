
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateQualityRuleComponent } from './create.component';
import { QualityRuleService } from '../../../services/QualityRule.service';
import { Router } from '@angular/router';

describe('CreateQualityRuleComponent', () => {
  let component: CreateQualityRuleComponent;
  let fixture: ComponentFixture<CreateQualityRuleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateQualityRuleComponent
      ],
      providers: [
        QualityRuleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateQualityRuleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});