
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBonusPlanComponent } from './index.component';
import { BonusPlanService } from '../../../services/BonusPlan.service';

describe('IndexBonusPlanComponent', () => {
  let component: IndexBonusPlanComponent;
  let fixture: ComponentFixture<IndexBonusPlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBonusPlanComponent
      ],
      providers: [
        BonusPlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBonusPlanComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});