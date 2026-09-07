
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCreditorComponent } from './index.component';
import { CreditorService } from '../../../services/Creditor.service';

describe('IndexCreditorComponent', () => {
  let component: IndexCreditorComponent;
  let fixture: ComponentFixture<IndexCreditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCreditorComponent
      ],
      providers: [
        CreditorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCreditorComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});