
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEnterpriseComponent } from './index.component';
import { EnterpriseService } from '../../../services/Enterprise.service';

describe('IndexEnterpriseComponent', () => {
  let component: IndexEnterpriseComponent;
  let fixture: ComponentFixture<IndexEnterpriseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEnterpriseComponent
      ],
      providers: [
        EnterpriseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEnterpriseComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});