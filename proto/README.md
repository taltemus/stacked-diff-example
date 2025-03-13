# Table of Contents

- [cart.v1](#cart-v1)
  - Services
    - [cart.v1.CartActor](#cart-v1-cartactor)
      - [Workflows](#cart-v1-cartactor-workflows)
        - [cart.v1.CartActor.StartCart](#cart-v1-cartactor-startcart-workflow)
      - [Signals](#cart-v1-cartactor-signals)
        - [cart.v1.CartActor.CancelCartSignal](#cart-v1-cartactor-cancelcartsignal-signal)
      - [Activities](#cart-v1-cartactor-activities)
        - [cart.v1.CartActor.ReserveTicketsActivity](#cart-v1-cartactor-reserveticketsactivity-activity)
        - [cart.v1.CartActor.RemoveTicketsActivity](#cart-v1-cartactor-removeticketsactivity-activity)
        - [cart.v1.CartActor.CheckoutCartActivity](#cart-v1-cartactor-checkoutcartactivity-activity)
  - Messages
    - [cart.v1.CancelCartSignalRequest](#cart-v1-cancelcartsignalrequest)
    - [cart.v1.CartRequest](#cart-v1-cartrequest)
    - [cart.v1.CartResponse](#cart-v1-cartresponse)
    - [cart.v1.CheckoutCartInput](#cart-v1-checkoutcartinput)
    - [cart.v1.CheckoutCartOutput](#cart-v1-checkoutcartoutput)
    - [cart.v1.RemoveTicketsInput](#cart-v1-removeticketsinput)
    - [cart.v1.RemoveTicketsOutput](#cart-v1-removeticketsoutput)
    - [cart.v1.ReserveTicketsInput](#cart-v1-reserveticketsinput)
    - [cart.v1.ReserveTicketsOutput](#cart-v1-reserveticketsoutput)

<a name="cart-v1"></a>
# cart.v1

<a name="cart-v1-services"></a>
## Services

<a name="cart-v1-cartactor"></a>
## cart.v1.CartActor

<pre>
The CartActor defines the actor-style cart workflow, its activities, and signal.
</pre>

<a name="cart-v1-cartactor-workflows"></a>
### Workflows

---
<a name="cart-v1-cartactor-startcart-workflow"></a>
### cart.v1.CartActor.StartCart

<pre>
Starts the shopping cart workflow.
</pre>

**Input:** [cart.v1.CartRequest](#cart-v1-cartrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
Comma-separated ticket IDs.<br>

json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>

**Output:** [cart.v1.CartResponse](#cart-v1-cartresponse)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>confirmation</td>
<td>string</td>
<td><pre>
json_name: confirmation
go_name: Confirmation</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>id</td><td><pre><code>cart/${! uuid_v4() }</code></pre></td></tr>
<tr><td>id_reuse_policy</td><td><pre><code>WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED</code></pre></td></tr>
</table>   

<a name="cart-v1-cartactor-signals"></a>
### Signals

---
<a name="cart-v1-cartactor-cancelcartsignal-signal"></a>
### cart.v1.CartActor.CancelCartSignal

<pre>
CancelCartSignal allows external callers to cancel the cart reservation.
</pre>

**Input:** [cart.v1.CancelCartSignalRequest](#cart-v1-cancelcartsignalrequest)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>reason</td>
<td>string</td>
<td><pre>
json_name: reason
go_name: Reason</pre></td>
</tr>
</table>  

<a name="cart-v1-cartactor-activities"></a>
### Activities

---
<a name="cart-v1-cartactor-reserveticketsactivity-activity"></a>
### cart.v1.CartActor.ReserveTicketsActivity

<pre>
ReserveTickets is an activity that reserves tickets.
</pre>

**Input:** [cart.v1.ReserveTicketsInput](#cart-v1-reserveticketsinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>reservation_duration_seconds</td>
<td>int32</td>
<td><pre>
json_name: reservationDurationSeconds
go_name: ReservationDurationSeconds</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>

**Output:** [cart.v1.ReserveTicketsOutput](#cart-v1-reserveticketsoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>status</td>
<td>string</td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table> 

---
<a name="cart-v1-cartactor-removeticketsactivity-activity"></a>
### cart.v1.CartActor.RemoveTicketsActivity

<pre>
RemoveTickets is an activity that removes tickets from the cart after a timeout.
</pre>

**Input:** [cart.v1.RemoveTicketsInput](#cart-v1-removeticketsinput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>

**Output:** [cart.v1.RemoveTicketsOutput](#cart-v1-removeticketsoutput)

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>status</td>
<td>string</td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table> 

---
<a name="cart-v1-cartactor-checkoutcartactivity-activity"></a>
### cart.v1.CartActor.CheckoutCartActivity

<pre>
CheckoutCartActivity is an activity that removes tickets from the cart after a timeout.
</pre>

**Input:** [cart.v1.CheckoutCartInput](#cart-v1-checkoutcartinput)



**Output:** [cart.v1.CheckoutCartOutput](#cart-v1-checkoutcartoutput)



**Defaults:**

<table>
<tr><th>Name</th><th>Value</th></tr>
<tr><td>retry_policy.max_attempts</td><td>3</td></tr>
<tr><td>start_to_close_timeout</td><td>30 seconds</td></tr>
</table>   

<a name="cart-v1-messages"></a>
## Messages

<a name="cart-v1-cancelcartsignalrequest"></a>
### cart.v1.CancelCartSignalRequest

<pre>
Signal to cancel the cart.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>reason</td>
<td>string</td>
<td><pre>
json_name: reason
go_name: Reason</pre></td>
</tr>
</table>



<a name="cart-v1-cartrequest"></a>
### cart.v1.CartRequest

<pre>
Message definitions.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>customer_id</td>
<td>string</td>
<td><pre>
json_name: customerId
go_name: CustomerId</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
Comma-separated ticket IDs.<br>

json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>



<a name="cart-v1-cartresponse"></a>
### cart.v1.CartResponse

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>confirmation</td>
<td>string</td>
<td><pre>
json_name: confirmation
go_name: Confirmation</pre></td>
</tr>
</table>



<a name="cart-v1-checkoutcartinput"></a>
### cart.v1.CheckoutCartInput



<a name="cart-v1-checkoutcartoutput"></a>
### cart.v1.CheckoutCartOutput



<a name="cart-v1-removeticketsinput"></a>
### cart.v1.RemoveTicketsInput

<pre>
RemoveTickets messages.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>



<a name="cart-v1-removeticketsoutput"></a>
### cart.v1.RemoveTicketsOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>status</td>
<td>string</td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>



<a name="cart-v1-reserveticketsinput"></a>
### cart.v1.ReserveTicketsInput

<pre>
ReserveTickets messages.
</pre>

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>cart_id</td>
<td>string</td>
<td><pre>
json_name: cartId
go_name: CartId</pre></td>
</tr><tr>
<td>reservation_duration_seconds</td>
<td>int32</td>
<td><pre>
json_name: reservationDurationSeconds
go_name: ReservationDurationSeconds</pre></td>
</tr><tr>
<td>ticket_ids</td>
<td>string</td>
<td><pre>
json_name: ticketIds
go_name: TicketIds</pre></td>
</tr>
</table>



<a name="cart-v1-reserveticketsoutput"></a>
### cart.v1.ReserveTicketsOutput

<table>
<tr>
<th>Attribute</th>
<th>Type</th>
<th>Description</th>
</tr>
<tr>
<td>status</td>
<td>string</td>
<td><pre>
json_name: status
go_name: Status</pre></td>
</tr>
</table>

