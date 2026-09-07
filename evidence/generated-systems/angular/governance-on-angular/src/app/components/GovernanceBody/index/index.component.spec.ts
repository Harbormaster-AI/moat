
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexGovernanceBodyComponent } from './index.component';
import { GovernanceBodyService } from '../../../services/GovernanceBody.service';

describe('IndexGovernanceBodyComponent', () => {
  let component: IndexGovernanceBodyComponent;
  let fixture: ComponentFixture<IndexGovernanceBodyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexGovernanceBodyComponent
      ],
      providers: [
        GovernanceBodyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexGovernanceBodyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});