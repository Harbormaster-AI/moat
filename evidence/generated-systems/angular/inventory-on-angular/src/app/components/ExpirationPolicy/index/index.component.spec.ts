
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexExpirationPolicyComponent } from './index.component';
import { ExpirationPolicyService } from '../../../services/ExpirationPolicy.service';

describe('IndexExpirationPolicyComponent', () => {
  let component: IndexExpirationPolicyComponent;
  let fixture: ComponentFixture<IndexExpirationPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexExpirationPolicyComponent
      ],
      providers: [
        ExpirationPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexExpirationPolicyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});