
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexUsageLimitComponent } from './index.component';
import { UsageLimitService } from '../../../services/UsageLimit.service';

describe('IndexUsageLimitComponent', () => {
  let component: IndexUsageLimitComponent;
  let fixture: ComponentFixture<IndexUsageLimitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexUsageLimitComponent
      ],
      providers: [
        UsageLimitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexUsageLimitComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});