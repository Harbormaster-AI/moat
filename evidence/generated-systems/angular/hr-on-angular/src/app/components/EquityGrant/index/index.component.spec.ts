
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEquityGrantComponent } from './index.component';
import { EquityGrantService } from '../../../services/EquityGrant.service';

describe('IndexEquityGrantComponent', () => {
  let component: IndexEquityGrantComponent;
  let fixture: ComponentFixture<IndexEquityGrantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEquityGrantComponent
      ],
      providers: [
        EquityGrantService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEquityGrantComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});