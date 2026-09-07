
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexUnderwritingDecisionComponent } from './index.component';
import { UnderwritingDecisionService } from '../../../services/UnderwritingDecision.service';

describe('IndexUnderwritingDecisionComponent', () => {
  let component: IndexUnderwritingDecisionComponent;
  let fixture: ComponentFixture<IndexUnderwritingDecisionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexUnderwritingDecisionComponent
      ],
      providers: [
        UnderwritingDecisionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexUnderwritingDecisionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});